Setup
  $ . ${TESTDIR}/../setup.sh
  $ . ${TESTDIR}/setup.sh $(pwd)

# Run the build task with --dry flag and cut up the logs into separate files by empty lines
# https://stackoverflow.com/a/33297878/986415
  $ ${TURBO} run build --dry |  awk -v RS= '{print > ("tmp-" NR ".txt")}'

# The first part of the file is Packages in Scope
  $ cat tmp-1.txt
  Packages in Scope
  Name    Path             
  another packages/another 
  my-app  apps/my-app      
  util    packages/util    

# Part 2 of the logs are Global Hash INputs
  $ cat tmp-2.txt
  Global Hash Inputs
    Global Files               = 1
    External Dependencies Hash = ccab0b28617f1f56
    Global Cache Key           = Buffalo buffalo Buffalo buffalo buffalo buffalo Buffalo buffalo
    Root pipeline              = {"//#something":{"outputs":[],"cache":true,"dependsOn":[],"inputs":[],"outputMode":"full","passthroughEnv":[],"env":[],"persistent":false},"build":{"outputs":[],"cache":true,"dependsOn":[],"inputs":[],"outputMode":"full","passthroughEnv":[],"env":["NODE_ENV"],"persistent":false},"maybefails":{"outputs":[],"cache":true,"dependsOn":[],"inputs":[],"outputMode":"full","passthroughEnv":[],"env":[],"persistent":false},"my-app#build":{"outputs":["apple.json","banana.txt"],"cache":true,"dependsOn":[],"inputs":[],"outputMode":"full","passthroughEnv":[],"env":[],"persistent":false},"something":{"outputs":[],"cache":true,"dependsOn":[],"inputs":[],"outputMode":"full","passthroughEnv":[],"env":[],"persistent":false}}

# Part 3 are Tasks to Run, and we have to validate each task separately
  $ cat tmp-3.txt | grep "my-app#build" -A 15
  my-app#build
    Task                             = build                                                                                                                                               
    Package                          = my-app                                                                                                                                              
    Hash                             = 2f192ed93e20f940                                                                                                                                    
    Cached (Local)                   = false                                                                                                                                               
    Cached (Remote)                  = false                                                                                                                                               
    Directory                        = apps/my-app                                                                                                                                         
    Command                          = echo 'building'                                                                                                                                     
    Outputs                          = apple.json, banana.txt                                                                                                                              
    Log File                         = apps/my-app/.turbo/turbo-build.log                                                                                                                  
    Dependencies                     =                                                                                                                                                     
    Dependendents                    =                                                                                                                                                     
    Inputs Files Considered          = 1                                                                                                                                                   
    Configured Environment Variables =                                                                                                                                                     
    Inferred Environment Variables   =                                                                                                                                                     
    Global Environment Variables     = SOME_ENV_VAR=, VERCEL_ANALYTICS_ID=                                                                                                                 

  $ cat tmp-3.txt | grep "util#build" -A 15
  util#build
    Task                             = build                                                                                                                                
    Package                          = util                                                                                                                                 
    Hash                             = af2ba2d52192ee45                                                                                                                     
    Cached (Local)                   = false                                                                                                                                
    Cached (Remote)                  = false                                                                                                                                
    Directory                        = packages/util                                                                                                                        
    Command                          = echo 'building'                                                                                                                      
    Outputs                          =                                                                                                                                      
    Log File                         = packages/util/.turbo/turbo-build.log                                                                                                 
    Dependencies                     =                                                                                                                                      
    Dependendents                    =                                                                                                                                      
    Inputs Files Considered          = 1                                                                                                                                    
    Configured Environment Variables = NODE_ENV=                                                                                                                            
    Inferred Environment Variables   =                                                                                                                                      
    Global Environment Variables     = SOME_ENV_VAR=, VERCEL_ANALYTICS_ID=                                                                                                  

# Run the task with NODE_ENV set and see it in summary. Use util package so it's just one package
  $ NODE_ENV=banana ${TURBO} run build --dry --filter=util | grep "Environment Variables"
    Configured Environment Variables = NODE_ENV=b493d48364afe44d11c0165cf470a4164d1e2609911ef998be868d46ade3de4e                                                            
    Inferred Environment Variables   =                                                                                                                                      
    Global Environment Variables     = SOME_ENV_VAR=, VERCEL_ANALYTICS_ID=                                                                                                  
