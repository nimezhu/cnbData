const {
    AppBar,
    Button,
    Typography,
    Card,
    CardContent,
    CardActions,
    CardActionArea,
    CardMedia,
    Icon,
    IconButton,
    Tabs,
    Tab,
    Table,
    TableCell,
    TableRow,
    TableHead,
    TableBody,
    Paper
} = window['material-ui'];


class ContactCard extends React.Component {
    render() {
        const {
            classes
        } = this.props
        return (<Card className={classes.card}>
        <CardContent>
          <Typography gutterBottom variant="h5" component="h2">
           Contact
          </Typography>
          <Typography component="p">
                Nucleome 
                Browser,
                is developed and maintained by <a href="mailto:zhuxp@cmu.edu">Xiaopeng Zhu</a> in <a href="http://www.cs.cmu.edu/~jianma/index.html" target="_blank">Jian Ma's Lab</a> in <a href="https://www.cmu.edu" target="_blank">Carnegie Mellon University</a>.
          </Typography>
        </CardContent>
        <CardActions>
        <Button size="small" color="primary">
          Learn More
        </Button>
        </CardActions>
        </Card>)
    }
}


/* Add Data Driven TrackDataCard */
class TrackCard extends React.Component {
    render() {
        const {
            classes,
            data,
            index
        } = this.props
        return (<Card className={classes.card}>
        <CardContent>
          <Typography gutterBottom variant="h5" component="h2">
            No.{data.name} {index}
          </Typography>
          <Typography component="p">
            TODO: Track Card Infomration
          </Typography>
        </CardContent>
        <CardActions>
        <Button size="small" color="primary">
          Learn More
        </Button>
        </CardActions>
        </Card>)

    }
}

class TrackDbList extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            genome: "hg38",
            dbs: []
        }
        var self = this
    }
    render() {
        const {
            classes,
            genome
        } = this.props
        var self = this
        const {
            dbs
        } = this.state
        fetch("/" + genome + "/ls").then(function(d) {
            return d.json()
        }).then(function(d) {
            self.setState({
                "genome": genome,
                "dbs": d
            })
        })
        return (<ul>
                {dbs.map((d,i)=>(<li>{d.dbname}
                    </li>))}
            </ul>)
    }
}


class TrackTable extends React.Component {
    render() {
        const {
            classes,
            rows
        } = this.props;

        return (
            <Paper className={classes.root}>
      <Table className={classes.table}>
        <TableHead>
          <TableRow>
            <TableCell>ID</TableCell>
            <TableCell align="right">Long Label</TableCell>
            <TableCell align="right">URI</TableCell>
            <TableCell align="right">Meta Link</TableCell>
            <TableCell align="right">Status</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {rows.map(row => (
            <TableRow key={row.id}>
              <TableCell component="th" scope="row">
                {row.id}
              </TableCell>
              <TableCell align="right">{row.longLabel}</TableCell>
              <TableCell align="right">{row.uri}</TableCell>
              <TableCell align="right">{row.metaLink}</TableCell>
              <TableCell align="right"></TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </Paper>
        );
    }
}
class TrackDb extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            tracks: []
        }
    }
    componentDidMount() {
        const {
            genome,
            dbname
        } = this.props
        var self = this
        fetch("/" + genome + "/" + dbname + "/list?attr=1").then(function(d) {
            return d.json()
        }).then(function(d) {
            self.setState({
                tracks: d
            })
        })

    }
    render() {
        const {
            classes
        } = this.props
        const {
            tracks
        } = this.state
        return (
            <div>
            <TrackTable classes={classes} rows={tracks}/>
        </div>
        )

    }
}

class TrackDbTabs extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            genome: "hg38",
            dbs: [],
            dbi: 0
        }
    }
    handleChange = (event, value) => {
        this.setState({
            dbi: value
        });
        console.log(this.state, value)
    };
    componentDidMount() {
        const {
            genome
        } = this.props
        var self = this
        fetch("/" + genome + "/ls").then(function(d) {
            return d.json()
        }).then(function(d) {
            self.setState({
                "genome": genome,
                "dbs": d
            })
        })
    }
    render() {
        const {
            classes,
            genome
        } = this.props
        const {
            dbi,
            dbs
        } = this.state
        return (
            <div>
        <AppBar position="static" color="default">
        
          <Tabs value={dbi} onChange={this.handleChange}  indicatorColor="primary"
          textColor="primary">
            {dbs.map((d,i)=>(
                <Tab label={d.dbname}/>
            ))}
          </Tabs>
        </AppBar>
            {dbs.map((d,i) => ( dbi === i && <TrackDb genome={genome} classes={classes} dbname={d.dbname} />))}
        </div>
        )
    }
}

/* Add Data Driven TrackDataCard */
class GenomeCard extends React.Component {
    render() {
        const {
            classes,
            data,
            index
        } = this.props
        return (<Card className={classes.card}>
        <CardContent>
          <Typography gutterBottom variant="h5" component="h2">
            {data}
          </Typography>
          <Typography component="p">
            <TrackDbTabs classes={classes} genome={data}/>
          </Typography>
        </CardContent>
        <CardActions>
        <Button size="small" color="primary">
          Learn More
        </Button>
        </CardActions>
        </Card>)

    }
}

class GenomeList extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            genomes: []
        }
        var self = this
        fetch("/genomes").then(function(d) {
            return d.json()
        }).then(function(d) {
            self.setState({
                "genomes": d
            })
        })
    }
    render() {
        const {
            classes
        } = this.props
        const {
            genomes
        } = this.state
        return (<div>
            {genomes.map((d,i) => (<GenomeCard classes={classes} data={d} index={i}/>) )}
            </div>)
    }

}
