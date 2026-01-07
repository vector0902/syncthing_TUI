
Besides using SYNCTHING_URL and SYNCTHING_API_KEY,
also support STHOME (syncthing home dir),
where program can find syncthing config file under this dir,
like syncthing_config_sample/config.xml
from which to parse SYNCTHING_URL and SYNCTHING_API_KEY.
example segment:

    <gui enabled="true" tls="false" sendBasicAuthPrompt="false">
        <address>127.0.0.1:8666</address>
        <metricsWithoutAuth>false</metricsWithoutAuth>
        <apikey>bJJfpeAN4WVPqTsePxvkdcMv59Mv2bPt</apikey>
        <theme>default</theme>
    </gui>


one thing to note about SYNCTHING_URL:
from this part:
<address>127.0.0.1:8666</address>
only get port,
then make SYNCTHING_URL=http://127.0.0.1:<port>

since sometimes <IP> may be something else.
