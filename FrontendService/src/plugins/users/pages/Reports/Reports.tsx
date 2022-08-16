import { useEffect, useState } from "react";
import { useQuery } from "react-query";
import { useReportsResult } from "../../hooks";
import { Card } from "../../imports";
import { useUserReportsService } from "../../services";
import { ReportsTable } from "./ReportsTable";
import { SearchReportForm } from "./SearchReportForm";

export const Reports = () => {

  const userReportsService = useUserReportsService();
  const [params, setParams] = useState<any>({});
  const { result, setResult } = useReportsResult();

  const { data: reports } = useQuery([result, params, userReportsService], () => userReportsService.search(params), { enabled: !result });

  useEffect(() => {
    if (!result) return;
    setResult(undefined);
  }, [result]);
  
  console.log(reports);


  return (
    <>
      <Card title="Reports">

        <SearchReportForm onSearch={setParams} />

        <ReportsTable reports={reports || []} />
      </Card>
    </>
  );  
}