package main

// Ejabberd - XMPP application server
// Homepage: https://www.ejabberd.im

import (
	"fmt"
	
	"os/exec"
)

func installEjabberd() {
	// Método 1: Descargar y extraer .tar.gz
	ejabberd_tar_url := "https://github.com/processone/ejabberd/archive/refs/tags/24.07.tar.gz"
	ejabberd_cmd_tar := exec.Command("curl", "-L", ejabberd_tar_url, "-o", "package.tar.gz")
	err := ejabberd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ejabberd_zip_url := "https://github.com/processone/ejabberd/archive/refs/tags/24.07.zip"
	ejabberd_cmd_zip := exec.Command("curl", "-L", ejabberd_zip_url, "-o", "package.zip")
	err = ejabberd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ejabberd_bin_url := "https://github.com/processone/ejabberd/archive/refs/tags/24.07.bin"
	ejabberd_cmd_bin := exec.Command("curl", "-L", ejabberd_bin_url, "-o", "binary.bin")
	err = ejabberd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ejabberd_src_url := "https://github.com/processone/ejabberd/archive/refs/tags/24.07.src.tar.gz"
	ejabberd_cmd_src := exec.Command("curl", "-L", ejabberd_src_url, "-o", "source.tar.gz")
	err = ejabberd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ejabberd_cmd_direct := exec.Command("./binary")
	err = ejabberd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: erlang")
	exec.Command("latte", "install", "erlang").Run()
	fmt.Println("Instalando dependencia: gd")
	exec.Command("latte", "install", "gd").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: linux-pam")
	exec.Command("latte", "install", "linux-pam").Run()
}
