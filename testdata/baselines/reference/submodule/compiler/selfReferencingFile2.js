//// [tests/cases/compiler/selfReferencingFile2.ts] ////

//// [selfReferencingFile2.ts]
///<reference path='../selfReferencingFile2.ts'/>

class selfReferencingFile2 {

}

//// [selfReferencingFile2.js]
//// [selfReferencingFile2.js]
class selfReferencingFile2 {
}
