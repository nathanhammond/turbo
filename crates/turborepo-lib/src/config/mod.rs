mod turbo;
mod turbo_config;

use thiserror::Error;
pub use turbo::{
    validate_extends, validate_no_package_task_syntax, RawTurboJSON, SpacesJson, TurboJson,
};
pub use turbo_config::{ConfigurationOptions, TurborepoConfigBuilder};

#[derive(Debug, Error)]
pub enum Error {
    #[error("default config path not found")]
    NoDefaultConfigPath,
    #[error(transparent)]
    PackageJson(#[from] crate::package_json::Error),
    #[error(
        "Could not find turbo.json. Follow directions at https://turbo.build/repo/docs to create \
         one"
    )]
    NoTurboJSON,
    #[error(
        "loginUrl is configured to \"{value}\", but cannot be a base URL. This happens in \
         situations like using a `data:` URL."
    )]
    LoginUrlCannotBeABase { value: String },
    #[error(transparent)]
    SerdeJson(#[from] serde_json::Error),
    #[error(transparent)]
    Io(#[from] std::io::Error),
    #[error(transparent)]
    Camino(#[from] camino::FromPathBufError),
    #[error(
        "Package tasks (<package>#<task>) are not allowed in single-package repositories: found \
         {task_id}"
    )]
    PackageTaskInSinglePackageMode { task_id: String },
    #[error(transparent)]
    Anyhow(#[from] anyhow::Error),
    #[error(
        "You specified \"{value}\" in the \"{key}\" key. You should not prefix your environment \
         variables with \"{env_pipeline_delimiter}\""
    )]
    InvalidEnvPrefix {
        value: String,
        key: String,
        env_pipeline_delimiter: &'static str,
    },
    #[error(transparent)]
    PathError(#[from] turbopath::PathError),
    #[error("\"{actual}\". Use \"{wanted}\" instead")]
    UnnecessaryPackageTaskSyntax { actual: String, wanted: String },
    #[error("You can only extend from the root workspace")]
    ExtendFromNonRoot,
    #[error("No \"extends\" key found")]
    NoExtends,
}
