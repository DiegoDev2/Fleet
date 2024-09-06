package main

// Snowflake - Pluggable Transport using WebRTC, inspired by Flashproxy
// Homepage: https://www.torproject.org

import (
	"fmt"
	
	"os/exec"
)

func installSnowflake() {
	// Método 1: Descargar y extraer .tar.gz
	snowflake_tar_url := "https://gitlab.torproject.org/tpo/anti-censorship/pluggable-transports/snowflake/-/archive/v2.9.2/snowflake-v2.9.2.tar.gz"
	snowflake_cmd_tar := exec.Command("curl", "-L", snowflake_tar_url, "-o", "package.tar.gz")
	err := snowflake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snowflake_zip_url := "https://gitlab.torproject.org/tpo/anti-censorship/pluggable-transports/snowflake/-/archive/v2.9.2/snowflake-v2.9.2.zip"
	snowflake_cmd_zip := exec.Command("curl", "-L", snowflake_zip_url, "-o", "package.zip")
	err = snowflake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snowflake_bin_url := "https://gitlab.torproject.org/tpo/anti-censorship/pluggable-transports/snowflake/-/archive/v2.9.2/snowflake-v2.9.2.bin"
	snowflake_cmd_bin := exec.Command("curl", "-L", snowflake_bin_url, "-o", "binary.bin")
	err = snowflake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snowflake_src_url := "https://gitlab.torproject.org/tpo/anti-censorship/pluggable-transports/snowflake/-/archive/v2.9.2/snowflake-v2.9.2.src.tar.gz"
	snowflake_cmd_src := exec.Command("curl", "-L", snowflake_src_url, "-o", "source.tar.gz")
	err = snowflake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snowflake_cmd_direct := exec.Command("./binary")
	err = snowflake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
