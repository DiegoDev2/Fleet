package main

// GnuGetopt - Command-line option parsing utility
// Homepage: https://github.com/util-linux/util-linux

import (
	"fmt"
	
	"os/exec"
)

func installGnuGetopt() {
	// Método 1: Descargar y extraer .tar.gz
	gnugetopt_tar_url := "https://mirrors.edge.kernel.org/pub/linux/utils/util-linux/v2.40/util-linux-2.40.2.tar.xz"
	gnugetopt_cmd_tar := exec.Command("curl", "-L", gnugetopt_tar_url, "-o", "package.tar.gz")
	err := gnugetopt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnugetopt_zip_url := "https://mirrors.edge.kernel.org/pub/linux/utils/util-linux/v2.40/util-linux-2.40.2.tar.xz"
	gnugetopt_cmd_zip := exec.Command("curl", "-L", gnugetopt_zip_url, "-o", "package.zip")
	err = gnugetopt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnugetopt_bin_url := "https://mirrors.edge.kernel.org/pub/linux/utils/util-linux/v2.40/util-linux-2.40.2.tar.xz"
	gnugetopt_cmd_bin := exec.Command("curl", "-L", gnugetopt_bin_url, "-o", "binary.bin")
	err = gnugetopt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnugetopt_src_url := "https://mirrors.edge.kernel.org/pub/linux/utils/util-linux/v2.40/util-linux-2.40.2.tar.xz"
	gnugetopt_cmd_src := exec.Command("curl", "-L", gnugetopt_src_url, "-o", "source.tar.gz")
	err = gnugetopt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnugetopt_cmd_direct := exec.Command("./binary")
	err = gnugetopt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoctor")
exec.Command("latte", "install", "asciidoctor")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
