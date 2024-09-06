package main

// OclIcd - OpenCL ICD loader
// Homepage: https://github.com/OCL-dev/ocl-icd/

import (
	"fmt"
	
	"os/exec"
)

func installOclIcd() {
	// Método 1: Descargar y extraer .tar.gz
	oclicd_tar_url := "https://github.com/OCL-dev/ocl-icd/archive/refs/tags/v2.3.2.tar.gz"
	oclicd_cmd_tar := exec.Command("curl", "-L", oclicd_tar_url, "-o", "package.tar.gz")
	err := oclicd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oclicd_zip_url := "https://github.com/OCL-dev/ocl-icd/archive/refs/tags/v2.3.2.zip"
	oclicd_cmd_zip := exec.Command("curl", "-L", oclicd_zip_url, "-o", "package.zip")
	err = oclicd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oclicd_bin_url := "https://github.com/OCL-dev/ocl-icd/archive/refs/tags/v2.3.2.bin"
	oclicd_cmd_bin := exec.Command("curl", "-L", oclicd_bin_url, "-o", "binary.bin")
	err = oclicd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oclicd_src_url := "https://github.com/OCL-dev/ocl-icd/archive/refs/tags/v2.3.2.src.tar.gz"
	oclicd_cmd_src := exec.Command("curl", "-L", oclicd_src_url, "-o", "source.tar.gz")
	err = oclicd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oclicd_cmd_direct := exec.Command("./binary")
	err = oclicd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
exec.Command("latte", "install", "asciidoc")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: opencl-headers")
exec.Command("latte", "install", "opencl-headers")
	fmt.Println("Instalando dependencia: xmlto")
exec.Command("latte", "install", "xmlto")
}
