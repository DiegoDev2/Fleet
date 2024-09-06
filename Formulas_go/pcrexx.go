package main

// Pcrexx - C++ wrapper for the Perl Compatible Regular Expressions
// Homepage: https://www.daemon.de/projects/pcrepp/

import (
	"fmt"
	
	"os/exec"
)

func installPcrexx() {
	// Método 1: Descargar y extraer .tar.gz
	pcrexx_tar_url := "https://www.daemon.de/idisk/Apps/pcre++/pcre++-0.9.5.tar.gz"
	pcrexx_cmd_tar := exec.Command("curl", "-L", pcrexx_tar_url, "-o", "package.tar.gz")
	err := pcrexx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pcrexx_zip_url := "https://www.daemon.de/idisk/Apps/pcre++/pcre++-0.9.5.zip"
	pcrexx_cmd_zip := exec.Command("curl", "-L", pcrexx_zip_url, "-o", "package.zip")
	err = pcrexx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pcrexx_bin_url := "https://www.daemon.de/idisk/Apps/pcre++/pcre++-0.9.5.bin"
	pcrexx_cmd_bin := exec.Command("curl", "-L", pcrexx_bin_url, "-o", "binary.bin")
	err = pcrexx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pcrexx_src_url := "https://www.daemon.de/idisk/Apps/pcre++/pcre++-0.9.5.src.tar.gz"
	pcrexx_cmd_src := exec.Command("curl", "-L", pcrexx_src_url, "-o", "source.tar.gz")
	err = pcrexx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pcrexx_cmd_direct := exec.Command("./binary")
	err = pcrexx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pcre")
exec.Command("latte", "install", "pcre")
}
