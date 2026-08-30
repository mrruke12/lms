# Versioning
#### Problem
What if the lesson was already completed by some students but the teacher decides to change it? In this scenario score stays immutable, however the content of the lesson is altered therefore user won't be able to see the old version of the lesson and his answers properly. 

#### Solution
1. Define a lifecycle of the lesson (draft-edit-published-archived statuses);
2. Define version field in the Lesson model;
3. Define a table to store lesson version snapshots; 
4. When the lesson is published do:
	1. Compare actual elements with the latest revisions;
	2. If an element is changed since the latest revision or is new then create new revision of this element;
	3. Take a snapshot of the lesson (update the version in the lesson table, and then write ids of all revisions with the latest created_at with the reference to this lesson version).
5. Now bound students completions to the lesson version as well so even if the lesson is changed student still can access the original view.
