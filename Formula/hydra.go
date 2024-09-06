package main

// Hydra - Network logon cracker which supports many services
// Homepage: https://github.com/vanhauser-thc/thc-hydra

import (
	"fmt"
	
	"os/exec"
)

func installHydra() {
	// Método 1: Descargar y extraer .tar.gz
	hydra_tar_url := "https://github.com/vanhauser-thc/thc-hydra/archive/refs/tags/v9.5.tar.gz"
	hydra_cmd_tar := exec.Command("curl", "-L", hydra_tar_url, "-o", "package.tar.gz")
	err := hydra_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hydra_zip_url := "https://github.com/vanhauser-thc/thc-hydra/archive/refs/tags/v9.5.zip"
	hydra_cmd_zip := exec.Command("curl", "-L", hydra_zip_url, "-o", "package.zip")
	err = hydra_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hydra_bin_url := "https://github.com/vanhauser-thc/thc-hydra/archive/refs/tags/v9.5.bin"
	hydra_cmd_bin := exec.Command("curl", "-L", hydra_bin_url, "-o", "binary.bin")
	err = hydra_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hydra_src_url := "https://github.com/vanhauser-thc/thc-hydra/archive/refs/tags/v9.5.src.tar.gz"
	hydra_cmd_src := exec.Command("curl", "-L", hydra_src_url, "-o", "source.tar.gz")
	err = hydra_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hydra_cmd_direct := exec.Command("./binary")
	err = hydra_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libssh")
	exec.Command("latte", "install", "libssh").Run()
	fmt.Println("Instalando dependencia: mysql-client")
	exec.Command("latte", "install", "mysql-client").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
