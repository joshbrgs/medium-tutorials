import type { MetaFunction } from "@remix-run/node";
import { useEffect, useState } from "react";
import { Form } from "@remix-run/react";
import { Switch } from "@/components/ui/switch";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { FeatureFlag } from "~/types/feature-flags";


export const meta: MetaFunction = () => {
  return [
    { title: "New Remix App" },
    { name: "description", content: "Welcome to Remix!" },
  ];
};

export default function FeatureFlagsPage() {
  const [flags, setFlags] = useState<FeatureFlag[]>([]);
  const [newFlagName, setNewFlagName] = useState("");

  useEffect(() => {
    fetch("http://localhost:3001/v1/flags")
      .then((res) => res.json())
      .then((data) => setFlags(data));
  }, []);

  const handleCreate = async (e: React.FormEvent) => {
    e.preventDefault();

    await fetch("http://localhost:3001/v1/flags", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        flag: newFlagName,
        variations: {
          default_var: false,
          false_var: false,
          true_var: true,
        },
        defaultRule: {
          percentage: {
            false_var: 100,
            true_var: 0,
          },
        },
      }),
    });

    setNewFlagName("");
    const res = await fetch("http://localhost:3001/v1/flags");
    setFlags(await res.json());
  };

  const handleToggle = async (flag: FeatureFlag) => {
    const newSplit = flag.defaultRule.percentage.true_var === 100
      ? { false_var: 100, true_var: 0 }
      : { false_var: 0, true_var: 100 };

    await fetch(`http://localhost:3001/v1/flags/${flag.id}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ defaultRule: { percentage: newSplit } }),
    });

    const res = await fetch("http://localhost:3001/v1/flags");
    setFlags(await res.json());
  };

  return (
    <div className="max-w-4xl mx-auto mt-10 space-y-10">
      <Card>
        <CardHeader>
          <CardTitle>Create New Feature Flag</CardTitle>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleCreate} className="flex gap-4">
            <div className="grid gap-2 w-full">
              <Label htmlFor="name">Flag Name</Label>
              <Input
                id="name"
                value={newFlagName}
                onChange={(e) => setNewFlagName(e.target.value)}
                placeholder="e.g. new-admin-ui"
                required
              />
            </div>
            <div className="flex items-end">
              <Button type="submit">Create</Button>
            </div>
          </form>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Feature Flags</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            {flags && flags.map((flag) => (
              <div key={flag.id} className="flex items-center justify-between border-b pb-2">
                <span className="text-sm font-medium">{flag.flag}</span>
                <Switch
                  checked={flag.defaultRule.percentage.true_var === 100}
                  onCheckedChange={() => handleToggle(flag)}
                />
              </div>
            ))}
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
