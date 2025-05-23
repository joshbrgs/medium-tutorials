import { OpenFeature, OpenFeatureProvider } from "@openfeature/react-sdk";
import { GoFeatureFlagWebProvider } from "@openfeature/go-feature-flag-web-provider";

const goFeatureFlagWebProvider = new GoFeatureFlagWebProvider({
  endpoint: "http://localhost:1031"
});

// Set the initial context for your evaluations
OpenFeature.setContext({
  targetingKey: "user-1",
  admin: false
});

// Instantiate and set our provider (be sure this only happens once)!
// Note: there's no need to await its initialization, the React SDK handles re-rendering and suspense for you!
OpenFeature.setProvider(goFeatureFlagWebProvider);

// Enclose your content in the configured provider
export function OFWrapper({ children }: { children: React.ReactNode }) {
  return (
    <OpenFeatureProvider>
      {children}
    </OpenFeatureProvider>
  );
}
