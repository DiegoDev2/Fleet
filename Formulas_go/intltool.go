package main

// Intltool - String tool
// Homepage: https://wiki.freedesktop.org/www/Software/intltool

import (
	"fmt"
	
	"os/exec"
)

func installIntltool() {
	// Método 1: Descargar y extraer .tar.gz
	intltool_tar_url := "https://launchpad.net/intltool/trunk/0.51.0/+download/intltool-0.51.0.tar.gz"
	intltool_cmd_tar := exec.Command("curl", "-L", intltool_tar_url, "-o", "package.tar.gz")
	err := intltool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	intltool_zip_url := "https://launchpad.net/intltool/trunk/0.51.0/+download/intltool-0.51.0.zip"
	intltool_cmd_zip := exec.Command("curl", "-L", intltool_zip_url, "-o", "package.zip")
	err = intltool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	intltool_bin_url := "https://launchpad.net/intltool/trunk/0.51.0/+download/intltool-0.51.0.bin"
	intltool_cmd_bin := exec.Command("curl", "-L", intltool_bin_url, "-o", "binary.bin")
	err = intltool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	intltool_src_url := "https://launchpad.net/intltool/trunk/0.51.0/+download/intltool-0.51.0.src.tar.gz"
	intltool_cmd_src := exec.Command("curl", "-L", intltool_src_url, "-o", "source.tar.gz")
	err = intltool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	intltool_cmd_direct := exec.Command("./binary")
	err = intltool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: expat")
exec.Command("latte", "install", "expat")
	fmt.Println("Instalando dependencia: perl-xml-parser")
exec.Command("latte", "install", "perl-xml-parser")
}
