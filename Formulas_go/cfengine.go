package main

// Cfengine - Help manage and understand IT infrastructure
// Homepage: https://cfengine.com/

import (
	"fmt"
	
	"os/exec"
)

func installCfengine() {
	// Método 1: Descargar y extraer .tar.gz
	cfengine_tar_url := "https://cfengine-package-repos.s3.amazonaws.com/tarballs/cfengine-community-3.24.0.tar.gz"
	cfengine_cmd_tar := exec.Command("curl", "-L", cfengine_tar_url, "-o", "package.tar.gz")
	err := cfengine_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cfengine_zip_url := "https://cfengine-package-repos.s3.amazonaws.com/tarballs/cfengine-community-3.24.0.zip"
	cfengine_cmd_zip := exec.Command("curl", "-L", cfengine_zip_url, "-o", "package.zip")
	err = cfengine_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cfengine_bin_url := "https://cfengine-package-repos.s3.amazonaws.com/tarballs/cfengine-community-3.24.0.bin"
	cfengine_cmd_bin := exec.Command("curl", "-L", cfengine_bin_url, "-o", "binary.bin")
	err = cfengine_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cfengine_src_url := "https://cfengine-package-repos.s3.amazonaws.com/tarballs/cfengine-community-3.24.0.src.tar.gz"
	cfengine_cmd_src := exec.Command("curl", "-L", cfengine_src_url, "-o", "source.tar.gz")
	err = cfengine_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cfengine_cmd_direct := exec.Command("./binary")
	err = cfengine_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: lmdb")
exec.Command("latte", "install", "lmdb")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: linux-pam")
exec.Command("latte", "install", "linux-pam")
}
