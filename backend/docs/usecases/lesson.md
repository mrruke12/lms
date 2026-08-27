# Lesson completion
**Actor**: student.
**Prerequisites**: logged in, has access to the lesson, hasn't completed it yet. 
**Trigger**: student begins the lesson.
**Result**: the lesson completes and student earns a grade.

#### Main scenario
1. Student begins the lesson;
2. Student aknowledges the information and answers the questions;
3. Student ends the lesson;
4. Student earns a grade and lesson completes.

#### Lesson status change to inactive with active sessions
1. Student begins the lesson;
2. Teacher sets lesson's status to different from "published";
3. Student aknowledges the information and answers the questions;
4. System's polling catches that the lesson is not published;
5. Student recieves a message that the lesson is unaviable now;
6. System saves the student's progress to use it when the lesson is aviable.
