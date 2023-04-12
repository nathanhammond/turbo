Setup
  $ . ${TESTDIR}/../_helpers/setup.sh
  $ . ${TESTDIR}/../_helpers/setup_monorepo.sh $(pwd) basic_monorepo

Baseline hashes
  $ ${TURBO} build --dry=json | jq -r '[.tasks[] | { taskId, hash, envMode, passthrough: .environmentVariables.passthrough, globalPassthrough: .environmentVariables.globalPassthrough }] | sort_by(.taskId)'
  [
    {
      "taskId": "another#build",
      "hash": "4147d91ceecf2174",
      "envMode": "Infer",
      "passthrough": null,
      "globalPassthrough": null
    },
    {
      "taskId": "my-app#build",
      "hash": "2f192ed93e20f940",
      "envMode": "Infer",
      "passthrough": null,
      "globalPassthrough": null
    },
    {
      "taskId": "util#build",
      "hash": "af2ba2d52192ee45",
      "envMode": "Infer",
      "passthrough": null,
      "globalPassthrough": null
    }
  ]

Add global passthrough
  $ cp "$TESTDIR/fixture-configs/set_difference_1_add_global.json" "$(pwd)/turbo.json" && git commit -am "no comment" --quiet
  $ ${TURBO} build --dry=json | jq -r '[.tasks[] | { taskId, hash, envMode, passthrough: .environmentVariables.passthrough, globalPassthrough: .environmentVariables.globalPassthrough }] | sort_by(.taskId)'
  [
    {
      "taskId": "another#build",
      "hash": "ea39568e9498f0cc",
      "envMode": "Strict",
      "passthrough": null,
      "globalPassthrough": [
        "DUPLICATED="
      ]
    },
    {
      "taskId": "my-app#build",
      "hash": "9e0607b00ae10739",
      "envMode": "Strict",
      "passthrough": null,
      "globalPassthrough": [
        "DUPLICATED="
      ]
    },
    {
      "taskId": "util#build",
      "hash": "13601ad4311021f9",
      "envMode": "Strict",
      "passthrough": null,
      "globalPassthrough": [
        "DUPLICATED="
      ]
    }
  ]

Skip below tests. They rely on global hash not including pipeline.
  $ exit 80

Add duplicate local passthrough
  $ cp "$TESTDIR/fixture-configs/set_difference_2_add_global_to_local.json" "$(pwd)/turbo.json" && git commit -am "no comment" --quiet
  $ ${TURBO} build --dry=json | jq -r '[.tasks[] | { taskId, hash, envMode, passthrough: .environmentVariables.passthrough, globalPassthrough: .environmentVariables.globalPassthrough }] | sort_by(.taskId)'
  [
    {
      "taskId": "another#build",
      "hash": "ea39568e9498f0cc",
      "envMode": "Strict",
      "passthrough": [
        "DUPLICATED="
      ],
      "globalPassthrough": [
        "DUPLICATED="
      ]
    },
    {
      "taskId": "my-app#build",
      "hash": "9e0607b00ae10739",
      "envMode": "Strict",
      "passthrough": [
        "DUPLICATED="
      ],
      "globalPassthrough": [
        "DUPLICATED="
      ]
    },
    {
      "taskId": "util#build",
      "hash": "13601ad4311021f9",
      "envMode": "Strict",
      "passthrough": [
        "DUPLICATED="
      ],
      "globalPassthrough": [
        "DUPLICATED="
      ]
    }
  ]
  
Add unique local passthrough
  $ cp "$TESTDIR/fixture-configs/set_difference_3_add_unique_local.json" "$(pwd)/turbo.json" && git commit -am "no comment" --quiet
  $ ${TURBO} build --dry=json | jq -r '[.tasks[] | { taskId, hash, envMode, passthrough: .environmentVariables.passthrough, globalPassthrough: .environmentVariables.globalPassthrough }] | sort_by(.taskId)'
  [
    {
      "taskId": "another#build",
      "hash": "UNKNOWN",
      "envMode": "Strict",
      "passthrough": [
        "DUPLICATED=",
        "UNIQUE="
      ],
      "globalPassthrough": [
        "DUPLICATED="
      ]
    },
    {
      "taskId": "my-app#build",
      "hash": "UNKNOWN",
      "envMode": "Strict",
      "passthrough": [
        "DUPLICATED=",
        "UNIQUE="
      ],
      "globalPassthrough": [
        "DUPLICATED="
      ]
    },
    {
      "taskId": "util#build",
      "hash": "UNKNOWN",
      "envMode": "Strict",
      "passthrough": [
        "DUPLICATED=",
        "UNIQUE="
      ],
      "globalPassthrough": [
        "DUPLICATED="
      ]
    }
  ]
 
