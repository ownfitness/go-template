const { setup, teardown, setData } = require("../firestore-test-helper");
const { doc, getDoc, setDoc, collection, getDocs, deleteDoc, updateDoc } = require("firebase/firestore");
const { assertFails, assertSucceeds } = require("@firebase/rules-unit-testing");

const mockUser = "userId1";
const mockData = {
    "users/userId1": {
        name: "user",
        surname: "one"
    }
};

const mockData2 = {
    "users/userId2": {
        name: "user",
        surname: "two"
    }
};

describe("User rules", function() {
    this.timeout(5000);

    before(async function() {
        // runs before all tests in this file regardless where this line is defined.
        testEnv = await setup();
    });

    after(async () => {
        await teardown();
    });

    it("Can read user documents that are my own", async () => {
        await setData(mockData);
        const user = await testEnv.authenticatedContext(mockUser);
        await assertSucceeds(getDoc(doc(user.firestore(), "users/userId1")));
    });

    it("Can't read documents that aren't mine", async () => {
        await setData(mockData2);
        const user = await testEnv.authenticatedContext(mockUser);
        await assertFails(getDoc(doc(user.firestore(), "users/userId2")));
    });

    it("Can update user documents that are mine", async () => {
        await setData(mockData);
        const user = await testEnv.authenticatedContext(mockUser);
        await assertSucceeds(updateDoc(doc(user.firestore(), "users/userId1"), { name: "test-user" }));
    });

    it("Can not update user documents that aren't mine", async () => {
        await setData(mockData2);
        const user = await testEnv.authenticatedContext(mockUser);
        await assertFails(updateDoc(doc(user.firestore(), "users/userId2"), { name: "test-user" }));
    });

    it("Can not delete my own user document", async () => {
        await setData(mockData);
        const user = await testEnv.authenticatedContext(mockUser);
        await assertFails(deleteDoc(doc(user.firestore(), "users/userId1")));
    });

    it("Can not delete someone elses user document", async () => {
        await setData(mockData2);
        const user = await testEnv.authenticatedContext(mockUser);
        await assertFails(deleteDoc(doc(user.firestore(), "users/userId2")));
    });

    it("Can not delete someone elses user document when not logged in", async () => {
        await setData(mockData);
        const user = await testEnv.unauthenticatedContext();
        await assertFails(deleteDoc(doc(user.firestore(), "users/userId1")));
    });
});