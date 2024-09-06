package main

// Cppcms - Free High Performance Web Development Framework
// Homepage: http://cppcms.com/wikipp/en/page/main

import (
	"fmt"
	
	"os/exec"
)

func installCppcms() {
	// Método 1: Descargar y extraer .tar.gz
	cppcms_tar_url := "https://downloads.sourceforge.net/project/cppcms/cppcms/1.2.1/cppcms-1.2.1.tar.bz2"
	cppcms_cmd_tar := exec.Command("curl", "-L", cppcms_tar_url, "-o", "package.tar.gz")
	err := cppcms_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cppcms_zip_url := "https://downloads.sourceforge.net/project/cppcms/cppcms/1.2.1/cppcms-1.2.1.tar.bz2"
	cppcms_cmd_zip := exec.Command("curl", "-L", cppcms_zip_url, "-o", "package.zip")
	err = cppcms_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cppcms_bin_url := "https://downloads.sourceforge.net/project/cppcms/cppcms/1.2.1/cppcms-1.2.1.tar.bz2"
	cppcms_cmd_bin := exec.Command("curl", "-L", cppcms_bin_url, "-o", "binary.bin")
	err = cppcms_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cppcms_src_url := "https://downloads.sourceforge.net/project/cppcms/cppcms/1.2.1/cppcms-1.2.1.tar.bz2"
	cppcms_cmd_src := exec.Command("curl", "-L", cppcms_src_url, "-o", "source.tar.gz")
	err = cppcms_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cppcms_cmd_direct := exec.Command("./binary")
	err = cppcms_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
