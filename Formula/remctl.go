package main

// Remctl - Client/server application for remote execution of tasks
// Homepage: https://www.eyrie.org/~eagle/software/remctl/

import (
	"fmt"
	
	"os/exec"
)

func installRemctl() {
	// Método 1: Descargar y extraer .tar.gz
	remctl_tar_url := "https://archives.eyrie.org/software/kerberos/remctl-3.18.tar.xz"
	remctl_cmd_tar := exec.Command("curl", "-L", remctl_tar_url, "-o", "package.tar.gz")
	err := remctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	remctl_zip_url := "https://archives.eyrie.org/software/kerberos/remctl-3.18.tar.xz"
	remctl_cmd_zip := exec.Command("curl", "-L", remctl_zip_url, "-o", "package.zip")
	err = remctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	remctl_bin_url := "https://archives.eyrie.org/software/kerberos/remctl-3.18.tar.xz"
	remctl_cmd_bin := exec.Command("curl", "-L", remctl_bin_url, "-o", "binary.bin")
	err = remctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	remctl_src_url := "https://archives.eyrie.org/software/kerberos/remctl-3.18.tar.xz"
	remctl_cmd_src := exec.Command("curl", "-L", remctl_src_url, "-o", "source.tar.gz")
	err = remctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	remctl_cmd_direct := exec.Command("./binary")
	err = remctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
