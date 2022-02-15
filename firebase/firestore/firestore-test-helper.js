const { readFileSync } = require("fs");
const { doc, setDoc, getDoc } = require("firebase/firestore");
const {
    initializeTestEnvironment,
    RulesTestEnvironment,
} = require("@firebase/rules-unit-testing");

const { path } = require("path");

const projectId = `rules-spec-firestore-${Date.now()}`;
let testEnv;

module.exports.setup = async () => {
    // Get the db linked to the new firebase app that we creted
    if(!testEnv){
        testEnv = await initializeTestEnvironment({
            projectId: projectId,
            firestore: {
                rules: readFileSync(__dirname + "/firestore.rules", "utf8")
            },
        });
    }
    // return the initialised app for testing
    return testEnv;
};

module.exports.setData = async (data) => {
    // Write mock documents with test rules
    await testEnv.withSecurityRulesDisabled(async context => {
        // set as Var as firestore fails otherwise
        const firestore = context.firestore();
        if (data) {
            for (const key in data) {
                await setDoc(doc(firestore, key), data[key]);
                // console.log("JSON String ", JSON.stringify(await getDoc(doc(firestore, key))));
            }
        }

    });
    return testEnv;
}

module.exports.teardown = async () => {
    // Delete all apps currently running in the firebase simulated environment
    await testEnv.clearFirestore();
    console.log("data torndown");
};
