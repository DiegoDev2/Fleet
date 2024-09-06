package main

// Viennacl - Linear algebra library for many-core architectures and multi-core CPUs
// Homepage: https://viennacl.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installViennacl() {
	// Método 1: Descargar y extraer .tar.gz
	viennacl_tar_url := "https://downloads.sourceforge.net/project/viennacl/1.7.x/ViennaCL-1.7.1.tar.gz"
	viennacl_cmd_tar := exec.Command("curl", "-L", viennacl_tar_url, "-o", "package.tar.gz")
	err := viennacl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	viennacl_zip_url := "https://downloads.sourceforge.net/project/viennacl/1.7.x/ViennaCL-1.7.1.zip"
	viennacl_cmd_zip := exec.Command("curl", "-L", viennacl_zip_url, "-o", "package.zip")
	err = viennacl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	viennacl_bin_url := "https://downloads.sourceforge.net/project/viennacl/1.7.x/ViennaCL-1.7.1.bin"
	viennacl_cmd_bin := exec.Command("curl", "-L", viennacl_bin_url, "-o", "binary.bin")
	err = viennacl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	viennacl_src_url := "https://downloads.sourceforge.net/project/viennacl/1.7.x/ViennaCL-1.7.1.src.tar.gz"
	viennacl_cmd_src := exec.Command("curl", "-L", viennacl_src_url, "-o", "source.tar.gz")
	err = viennacl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	viennacl_cmd_direct := exec.Command("./binary")
	err = viennacl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: opencl-headers")
	exec.Command("latte", "install", "opencl-headers").Run()
	fmt.Println("Instalando dependencia: opencl-icd-loader")
	exec.Command("latte", "install", "opencl-icd-loader").Run()
	fmt.Println("Instalando dependencia: pocl")
	exec.Command("latte", "install", "pocl").Run()
}
