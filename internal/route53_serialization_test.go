package internal_test

import (
	"encoding/xml"
	"testing"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/gen/route53"
)

func TestRoute53RequestSerialization(t *testing.T) {
	r := &route53.ChangeResourceRecordSetsRequest{
		ChangeBatch: &route53.ChangeBatch{
			Comment: aws.String("hello"),
			Changes: []route53.Change{
				route53.Change{
					Action: aws.String("dance"),
					ResourceRecordSet: &route53.ResourceRecordSet{
						AliasTarget: &route53.AliasTarget{
							EvaluateTargetHealth: aws.False(),
						},
					},
				},
			},
		},
	}

	out, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	expected := `<ChangeResourceRecordSetsRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
  <ChangeBatch>
    <Change>
      <Action>dance</Action>
      <ResourceRecordSet>
        <AliasTarget>
          <EvaluateTargetHealth>false</EvaluateTargetHealth>
        </AliasTarget>
      </ResourceRecordSet>
    </Change>
    <Comment>hello</Comment>
  </ChangeBatch>
</ChangeResourceRecordSetsRequest>`

	if v, want := string(out), expected; v != want {
		t.Errorf("Was \n%s\n but expected \n%s", v, want)
	}
}
