const ProfessionalAccountSetup = (props: any) => {
  const {
    searchParams: { plan, next },
  } = props;

  return <div>Professional Account Setup: {JSON.stringify(props)}</div>;
};

export default ProfessionalAccountSetup;
