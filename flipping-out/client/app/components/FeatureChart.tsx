import {
  LineChart,
  Line,
  CartesianGrid,
  XAxis,
  YAxis,
  Tooltip,
} from "recharts";

type FeatureEvent = {
  timestamp: number;
  value: boolean;
};

interface Props {
  data: { timestamp: number; percentTrue: number }[];
}

export default function FeatureChart({ data }: Props) {
  return (
    <LineChart width={600} height={300} data={data}>
      <Line type="monotone" dataKey="percentTrue" stroke="#8884d8" />
      <CartesianGrid stroke="#ccc" />
      <XAxis dataKey="timestamp" />
      <YAxis domain={[0, 100]} unit="%" />
      <Tooltip />
    </LineChart>
  );
}
