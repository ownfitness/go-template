const { setup, teardown, setData } = require("../firestore-test-helper");
const { doc, getDoc, setDoc, collection, getDocs, deleteDoc } = require("firebase/firestore");
const { assertFails, assertSucceeds } = require("@firebase/rules-unit-testing");

const testUser = "testid";
let testEnv;

describe("Firestore rules", function() {
    this.timeout(5000);

    before(async function() {
        // runs before all tests in this file regardless where this line is defined.
        testEnv = await setup();
    });

    after(teardown);

    it("should deny a read to the private collection ", done => {
        const mockData = {
            "private/super-secret-document": {
                testData: "test"
            }
        };
        setData(mockData);
        const user = testEnv.authenticatedContext(testUser);
        assertFails(getDocs(collection(user.firestore(), "private")));
        done();
    });

    it("should allow a read to cities if logged in user reads it", done => {
        const mockData = {
            "cities/secret-city": {
                testData: "test"
            }
        };
        setData(mockData);
        const user = testEnv.authenticatedContext(testUser);
        assertSucceeds(getDocs(collection(user.firestore(), "cities")));
        done();
    });

    it("should deny a read to cities if a logged out user reads it", done => {
        const mockData = {
            "cities/secret-city": {
                testData: "test"
            }
        };

        setData(mockData);
        const user = testEnv.unauthenticatedContext();
        const firestore = user.firestore();
        assertFails(getDocs(collection(firestore, "cities")));
        assertFails(getDoc(doc(firestore, "cities/secret-city")));
        done();
    });

    it("should fail to delete cos it's not yours", done => {
        const mockData = {
            "posts/id1": {
                uid: "not_matching"
            }
        };

        const mockUser = "matchingUser";

        setData(mockData);
        const user = testEnv.authenticatedContext(mockUser);
        assertFails(deleteDoc(doc(user.firestore(), "posts/id1")));
        done();
    });

    it("should fail to delete cos you're not logged in", done => {
        const mockData = {
            "posts/id1": {
                ownerUID: "not_matching"
            }
        };

        setData(mockData);
        const user = testEnv.unauthenticatedContext();
        assertFails(deleteDoc(doc(user.firestore(), "posts/id1")));
        done();
    });

    it("should delete cos it's yours", done => {
        const mockData = {
            "posts/id2": {
                text: "bananas",
                uid: "matchingUser"
            }
        };

        setData(mockData);
        const user = testEnv.authenticatedContext("matchingUser");

        assertSucceeds(deleteDoc(doc(user.firestore(), "posts/id2")));

        const user2 = testEnv.authenticatedContext("notMatchingUser");
        assertFails(deleteDoc(doc(user2.firestore(), "posts/id2")));
        done();
    });

});