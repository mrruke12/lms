# Versioning
#### Problem
What happens if an author wants to change a lesson already completed or being in work by students? In first scenario nothing scary happens - students just see an updated lesson with their answers (where tasks didn't change) and the score **stays the same** (the score is immutable since the lesson is marked as completed by the student); but the second scenario may cause a problem if tasks changed (correct answer altered, task deleted, etc) so server can't handle answer check. 

#### Solution
1. Determine the lesson lifecycle (status such as draft, published, editing, archived);
2. Use versioning.
Versioning here solves the problem of 