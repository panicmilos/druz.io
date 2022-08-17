import { useEffect } from "react";
import { useQuery } from "react-query";
import { usePostsResult } from "../../hooks";
import { Card } from "../../imports";
import { usePostsService } from "../../services";
import { PostsList } from "./PostsList";

export const Posts = () => {


  const postsService = usePostsService();
  const { result, setResult } = usePostsResult();

  const { data: posts } = useQuery([postsService], () => postsService.fetchAll(), { enabled: !result });

  useEffect(() => {
    if (!result) return;
    setResult(undefined);
  }, [result]);
  
  console.log(posts);

  return (
    <>
      <Card title="Posts">

        <PostsList posts={posts || []} />
      </Card>

    </>
  ); 
  
}