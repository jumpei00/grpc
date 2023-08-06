import { getTodo } from "@/api/todo";
import Delete from "@/components/submit/delete";
import Update from "@/components/submit/update";

export default async function Page({ params }: { params: { id: string } }) {
    const data = await getTodo(params.id);

    return (
        <>
            <div>変更できます</div>
            <Update key={params.id} content={data.content} />
            <Delete />
        </>
    );
}
