export type FeatureFlag = {
  id: string;
  flag: string;
  variations: {
    default_var: boolean;
    false_var: boolean;
    true_var: boolean;
  };
  defaultRule: {
    percentage: {
      false_var: number;
      true_var: number;
    };
  };
};
