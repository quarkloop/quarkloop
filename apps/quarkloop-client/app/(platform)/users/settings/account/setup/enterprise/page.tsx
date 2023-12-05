const EnterpriseAccountSetup = (props: any) => {
  const {
    searchParams: { plan, next },
  } = props;

  return <div>Enterprise Account Setup: {JSON.stringify(props)}</div>;
};

export default EnterpriseAccountSetup;
