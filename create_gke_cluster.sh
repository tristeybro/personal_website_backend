while getopts c:z:n: flag
do
    case "${flag}" in
        c) cluster_name=${OPTARG};;
        z) cluster_zone=${OPTARG};;
        n) num_nodes=${OPTARG};;
    esac
done

response=$(gcloud container clusters describe $cluster_name --zone $cluster_zone || echo "ClusterNotFound")
if [[ $response = "ClusterNotFound" ]]
then
  echo "cluster not found, creating cluster"
  gcloud container clusters create $cluster_name --num-nodes=$num_nodes --zone=$cluster_zone
else
  echo "cluster already exists, skipping creating cluster"
  gcloud container clusters get-credentials $cluster_name --zone=$cluster_zone
fi
