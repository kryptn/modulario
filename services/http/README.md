# Data subproject

This is meant to isolate all of the db transactions within the app, regardless of deployment type or decided archetecture



### Post Decider Types:
1. Random:  Random
    * Needs only one decider model
2. Popularity: Steps based off of reddit points
    * Needs one decider model per link
    * visits with referrals will be parsed and will try to attach the reddit post to the modulario post
    * two reddit posts should be able to visit the link with different results
3. Altering
    * Based on points, do something different idk
    
    
### The Decider

post -> []decider

if only one, it's probably random, but the decider will know

If more than one, iterate through and build a closure to evaluate the decision


