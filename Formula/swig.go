package main

// Swig - Generate scripting interfaces to C/C++ code
// Homepage: https://www.swig.org/

import (
	"fmt"
	
	"os/exec"
)

func installSwig() {
	// Método 1: Descargar y extraer .tar.gz
	swig_tar_url := "https://downloads.sourceforge.net/project/swig/swig/swig-4.2.1/swig-4.2.1.tar.gz"
	swig_cmd_tar := exec.Command("curl", "-L", swig_tar_url, "-o", "package.tar.gz")
	err := swig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	swig_zip_url := "https://downloads.sourceforge.net/project/swig/swig/swig-4.2.1/swig-4.2.1.zip"
	swig_cmd_zip := exec.Command("curl", "-L", swig_zip_url, "-o", "package.zip")
	err = swig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	swig_bin_url := "https://downloads.sourceforge.net/project/swig/swig/swig-4.2.1/swig-4.2.1.bin"
	swig_cmd_bin := exec.Command("curl", "-L", swig_bin_url, "-o", "binary.bin")
	err = swig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	swig_src_url := "https://downloads.sourceforge.net/project/swig/swig/swig-4.2.1/swig-4.2.1.src.tar.gz"
	swig_cmd_src := exec.Command("curl", "-L", swig_src_url, "-o", "source.tar.gz")
	err = swig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	swig_cmd_direct := exec.Command("./binary")
	err = swig_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
