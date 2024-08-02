const env = import.meta.env.PROD ? "prod" : "local";

const BaseDomainProduction = "https://labeler-enbwmscxnq-as.a.run.app";

export const EnvNameConfig = {
  LOCAL: "local",
  PRODUCTION: "prod",
};

export const Configs = {
  [EnvNameConfig.LOCAL]: {
    API_BASE_URL: "https://labeler-enbwmscxnq-as.a.run.app",
  },
  [EnvNameConfig.PRODUCTION]: {
    API_BASE_URL: `https://${BaseDomainProduction}`,
  },
};

export const Config = Configs[env];
export default env;
