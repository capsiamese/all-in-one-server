import {Breadcrumb, BreadcrumbItem, BreadcrumbLink, Heading} from "@chakra-ui/react";
import {Link as RouteLink} from 'react-router-dom';

export default function HeadNav({routes}: { routes: object }) {
    return (
        <Heading>
            <Breadcrumb>
                {Object.entries(routes).map(([k, v]) => {
                    return (
                        <BreadcrumbItem key={k}>
                            <BreadcrumbLink as={RouteLink} to={v}>
                                {k}
                            </BreadcrumbLink>
                        </BreadcrumbItem>
                    )
                })}
            </Breadcrumb>
        </Heading>
    )
}
