import {Component} from "react";

function ResourcesList({results}) {
    return (
        <ul>
            {
                results.map(({resource, label}) =>
                    <li key={resource}>{label}</li>
                )
            }
        </ul>
    )
}
export default class SPARQLResults extends Component {
    render(){
        return ResourcesList(
            {results}
        )
    };
}

