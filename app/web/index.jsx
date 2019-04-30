const {
    Button,
    Menu,
    MenuItem,
    AppBar,
    Toolbar,
    Drawer,
    colors,
    createMuiTheme,
    CssBaseline,
    MuiThemeProvider,
    Typography,
    withStyles,
    Card,
    CardContent,
    CardActions,
    CardActionArea,
    CardMedia,
    Icon,
    IconButton
} = window['material-ui'];


const theme = createMuiTheme({
    palette: {
        primary: {
            light: colors.purple[300],
            main: colors.purple[500],
            dark: colors.purple[700],
        },
        secondary: {
            light: colors.green[300],
            main: colors.green[500],
            dark: colors.green[700],
        },
    },
    typography: {
        useNextVariants: true,
    },
});

const drawerWidth = 240;
const styles = theme => ({
    icon: {
        marginRight: theme.spacing.unit,
    },
   card: {
        minWidth: 275,
        maxWidth: 400,
        margin: 10,
        float: "left"
    },
    bullet: {
        display: 'inline-block',
        margin: '0 2px',
        transform: 'scale(1.8)',
    },

    title: {
        fontSize: 14,
    },
    pos: {
        marginBottom: 12,
    },
    media: {
        height: 0,
        paddingTop: '56.25%', // 16:9
    },
    actions: {
        display: 'flex',
    },
    expand: {
        transform: 'rotate(0deg)',
        marginLeft: 'auto',
        transition: theme.transitions.create('transform', {
            duration: theme.transitions.duration.shortest,
        }),
    },
    expandOpen: {
        transform: 'rotate(180deg)',
    },
    avatar: {
        backgroundColor: "#550000",
    },
    menuButton: {
        marginLeft: -18,
        marginRight: 10,
    },

    hide: {
        display: 'none',
    },
    drawer: {
        width: drawerWidth,
        flexShrink: 0,
    },
    drawerPaper: {
        width: drawerWidth,
    },
    drawerHeader: {
        display: 'flex',
        alignItems: 'center',
        padding: '0 8px',
        ...theme.mixins.toolbar,
        justifyContent: 'flex-end',
    },
    content: {
        flexGrow: 1,
        padding: theme.spacing.unit * 3,
        transition: theme.transitions.create('margin', {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.leavingScreen,
        }),
        marginLeft: -drawerWidth,
    },
    contentShift: {
        transition: theme.transitions.create('margin', {
            easing: theme.transitions.easing.easeOut,
            duration: theme.transitions.duration.enteringScreen,
        }),
        marginLeft: 0,
    },

});

/* Nav Bar Begin */
class SimpleAppBar extends React.Component {
    clickHandleFactory = (d) => {
        return function() {
            console.log("handle click", d)
            window.location.href = d + ".html"
        }
    };
    render() {
        var data = [{
                label: "Home",
                icon: "home",
                page: "index"
            },
            {
                label: "Admin",
                icon: "email",
                page: "admin"
            },
        ]
        var self = this;
        const {
            classes
        } = this.props
        var select = this.props.select
        data.forEach(function(d) {
            if (d.label == select) {
                d.color = "primary"
            }
        })
        return (
            <div>
      <AppBar position="static" color="default">
        <Toolbar className={classes.toolbar}>
            {data.map((d,i) => {
                console.log(d,i)
                return (<IconButton color={d.color || "inherit"} aria-label={d.label} onClick={self.clickHandleFactory(d.page)}> 
                                <Icon>{d.icon}</Icon> 
                        </IconButton>)
            })
            }
       </Toolbar>
      </AppBar>
    </div>
        );
    }
};
/* NavBar End */

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



class Index extends React.Component {
    state = {
        open: false,
    };

    handleDrawerOpen = () => {
        this.setState({
            open: true
        });
    };

    handleDrawerClose = () => {
        this.setState({
            open: false
        });
    };

    render() {
        const {
            classes,
            theme
        } = this.props;
        const {
            open
        } = this.state;

        console.log("index classes", classes)
        return (
        <MuiThemeProvider theme={theme}>
        <div className={classes.root}>
        <SimpleAppBar select="Home" classes={classes}>
        </SimpleAppBar>
        <ContactCard classes={classes}>
        </ContactCard>
       </div>

      </MuiThemeProvider>
        );
    }
}

const App = withStyles(styles)(Index);

ReactDOM.render(<App />, document.getElementById('root'));
