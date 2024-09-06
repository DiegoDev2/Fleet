package main

// Subversion - Version control system designed to be a better CVS
// Homepage: https://subversion.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installSubversion() {
	// Método 1: Descargar y extraer .tar.gz
	subversion_tar_url := "https://www.apache.org/dyn/closer.lua?path=subversion/subversion-1.14.3.tar.bz2"
	subversion_cmd_tar := exec.Command("curl", "-L", subversion_tar_url, "-o", "package.tar.gz")
	err := subversion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	subversion_zip_url := "https://www.apache.org/dyn/closer.lua?path=subversion/subversion-1.14.3.tar.bz2"
	subversion_cmd_zip := exec.Command("curl", "-L", subversion_zip_url, "-o", "package.zip")
	err = subversion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	subversion_bin_url := "https://www.apache.org/dyn/closer.lua?path=subversion/subversion-1.14.3.tar.bz2"
	subversion_cmd_bin := exec.Command("curl", "-L", subversion_bin_url, "-o", "binary.bin")
	err = subversion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	subversion_src_url := "https://www.apache.org/dyn/closer.lua?path=subversion/subversion-1.14.3.tar.bz2"
	subversion_cmd_src := exec.Command("curl", "-L", subversion_src_url, "-o", "source.tar.gz")
	err = subversion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	subversion_cmd_direct := exec.Command("./binary")
	err = subversion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: scons")
	exec.Command("latte", "install", "scons").Run()
	fmt.Println("Instalando dependencia: swig")
	exec.Command("latte", "install", "swig").Run()
	fmt.Println("Instalando dependencia: apr")
	exec.Command("latte", "install", "apr").Run()
	fmt.Println("Instalando dependencia: apr-util")
	exec.Command("latte", "install", "apr-util").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: utf8proc")
	exec.Command("latte", "install", "utf8proc").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
