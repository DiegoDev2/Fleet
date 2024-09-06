package main

// Scrollkeeper - Transitional package for scrollkeeper
// Homepage: https://scrollkeeper.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installScrollkeeper() {
	// Método 1: Descargar y extraer .tar.gz
	scrollkeeper_tar_url := "https://downloads.sourceforge.net/project/scrollkeeper/scrollkeeper/0.3.14/scrollkeeper-0.3.14.tar.gz"
	scrollkeeper_cmd_tar := exec.Command("curl", "-L", scrollkeeper_tar_url, "-o", "package.tar.gz")
	err := scrollkeeper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scrollkeeper_zip_url := "https://downloads.sourceforge.net/project/scrollkeeper/scrollkeeper/0.3.14/scrollkeeper-0.3.14.zip"
	scrollkeeper_cmd_zip := exec.Command("curl", "-L", scrollkeeper_zip_url, "-o", "package.zip")
	err = scrollkeeper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scrollkeeper_bin_url := "https://downloads.sourceforge.net/project/scrollkeeper/scrollkeeper/0.3.14/scrollkeeper-0.3.14.bin"
	scrollkeeper_cmd_bin := exec.Command("curl", "-L", scrollkeeper_bin_url, "-o", "binary.bin")
	err = scrollkeeper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scrollkeeper_src_url := "https://downloads.sourceforge.net/project/scrollkeeper/scrollkeeper/0.3.14/scrollkeeper-0.3.14.src.tar.gz"
	scrollkeeper_cmd_src := exec.Command("curl", "-L", scrollkeeper_src_url, "-o", "source.tar.gz")
	err = scrollkeeper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scrollkeeper_cmd_direct := exec.Command("./binary")
	err = scrollkeeper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docbook")
	exec.Command("latte", "install", "docbook").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: intltool")
	exec.Command("latte", "install", "intltool").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
