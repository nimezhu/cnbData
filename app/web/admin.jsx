const {
    CssBaseline,
    AppBar,
    Toolbar,
    Drawer,
    List,
    Icon,
    Typography,
    Divider,
    IconButton,
    ListItem,
    ListItemText,
    withStyles,
    createMuiTheme,
} = window['material-ui'];


const drawerWidth = 240;

const styles = theme => ({
    root: {
        display: 'flex',
    },
    appBar: {
        transition: theme.transitions.create(['margin', 'width'], {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.leavingScreen,
        }),
    },
    appBarShift: {
        width: `calc(100% - ${drawerWidth}px)`,
        marginLeft: drawerWidth,
        transition: theme.transitions.create(['margin', 'width'], {
            easing: theme.transitions.easing.easeOut,
            duration: theme.transitions.duration.enteringScreen,
        }),
    },
    menuButton: {
        marginLeft: 12,
        marginRight: 20,
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

class PersistentDrawerLeft extends React.Component {
    state = {
        open: true,
        nav:"Apps"
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
    handleNav = (d) => {
        var self = this
        return function(){
        console.log("todo nav",d)
        self.setState({
                nav:d
         })
         console.log(self.state)
        }
    }

    render() {
        const {
            classes,
            theme
        } = this.props;
        const {
            open,nav
        } = this.state;

        return (
            <div className={classes.root}>
        <CssBaseline />
        <AppBar
          position="fixed"
          className={classNames(classes.appBar, {
            [classes.appBarShift]: open,
          })}
        >
          <Toolbar disableGutters={!open}>
            <IconButton
              color="inherit"
              aria-label="Open drawer"
              onClick={this.handleDrawerOpen}
              className={classNames(classes.menuButton, open && classes.hide)}
            >
              <Icon>menu</Icon>
            </IconButton>
            <Typography variant="h6" color="inherit" noWrap>
                Nucleome Data Manager : {nav}
            </Typography>
          </Toolbar>
        </AppBar>
        <Drawer
          className={classes.drawer}
          variant="persistent"
          anchor="left"
          open={open}
          classes={{
            paper: classes.drawerPaper,
          }}
        >
          <div className={classes.drawerHeader}>
            <IconButton onClick={this.handleDrawerClose}>
              {theme.direction === 'ltr' ? <Icon>chevron_left</Icon> : <Icon>chevron_right </Icon>}
            </IconButton>
          </div>
          <Divider />
          <List>
            {[{icon:"apps",label:"Apps"},{icon:"cloud",label:"Data"},{icon:"vpn_key",label:"Admin"}].map((d, index) => (
              <ListItem button key={d.label} onClick = {this.handleNav(d.label)}>
                <Icon>{d.icon}</Icon>
                <ListItemText primary={d.label} />
              </ListItem>
            ))}
          </List>
          <Divider />
        </Drawer>
        <main
          className={classNames(classes.content, {
            [classes.contentShift]: open,
          })}
        >
          <div className={classes.drawerHeader} />
          {nav == "Data" && <div>
              <GenomeList classes={classes}/>
          </div>  
          }
          {nav == "Apps"?
          <Typography paragraph>
            Apps Paragraph
            <hr/>
            <div>
            <a href="https://vis.nucleome.org" target="_blank">Nucleome Browser</a>
            </div>
        </Typography>: null}
        {nav == "Admin"?
          <Typography paragraph>
             Admin
            <hr/>
            <div>
                Input Password ...
            </div>
            <ContactCard classes={classes}/>
        </Typography>: null}


        </main>
      </div>
        );
    }
}

PersistentDrawerLeft.propTypes = {
    classes: PropTypes.object.isRequired,
    theme: PropTypes.object.isRequired,
};

var App = withStyles(styles, {
    withTheme: true
})(PersistentDrawerLeft);
ReactDOM.render(<App />, document.getElementById('root'));
