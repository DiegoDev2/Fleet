package main

// Obfs4proxy - Pluggable transport proxy for Tor, implementing obfs4
// Homepage: https://gitlab.com/yawning/obfs4

import (
	"fmt"
	
	"os/exec"
)

func installObfs4proxy() {
	// Método 1: Descargar y extraer .tar.gz
	obfs4proxy_tar_url := "https://gitlab.com/yawning/obfs4/-/archive/obfs4proxy-0.0.14/obfs4-obfs4proxy-0.0.14.tar.gz"
	obfs4proxy_cmd_tar := exec.Command("curl", "-L", obfs4proxy_tar_url, "-o", "package.tar.gz")
	err := obfs4proxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	obfs4proxy_zip_url := "https://gitlab.com/yawning/obfs4/-/archive/obfs4proxy-0.0.14/obfs4-obfs4proxy-0.0.14.zip"
	obfs4proxy_cmd_zip := exec.Command("curl", "-L", obfs4proxy_zip_url, "-o", "package.zip")
	err = obfs4proxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	obfs4proxy_bin_url := "https://gitlab.com/yawning/obfs4/-/archive/obfs4proxy-0.0.14/obfs4-obfs4proxy-0.0.14.bin"
	obfs4proxy_cmd_bin := exec.Command("curl", "-L", obfs4proxy_bin_url, "-o", "binary.bin")
	err = obfs4proxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	obfs4proxy_src_url := "https://gitlab.com/yawning/obfs4/-/archive/obfs4proxy-0.0.14/obfs4-obfs4proxy-0.0.14.src.tar.gz"
	obfs4proxy_cmd_src := exec.Command("curl", "-L", obfs4proxy_src_url, "-o", "source.tar.gz")
	err = obfs4proxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	obfs4proxy_cmd_direct := exec.Command("./binary")
	err = obfs4proxy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
