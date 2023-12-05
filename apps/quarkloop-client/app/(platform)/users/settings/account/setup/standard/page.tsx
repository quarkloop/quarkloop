const StandardAccountSetup = (props: any) => {
  const {
    searchParams: { plan, next },
  } = props;

  return <div>Standard Account Setup: {JSON.stringify(props)}</div>;
};

export default StandardAccountSetup;
