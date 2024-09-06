package main

// Cf2tf - Cloudformation templates to Terraform HCL converter
// Homepage: https://github.com/DontShaveTheYak/cf2tf

import (
	"fmt"
	
	"os/exec"
)

func installCf2tf() {
	// Método 1: Descargar y extraer .tar.gz
	cf2tf_tar_url := "https://files.pythonhosted.org/packages/52/00/94c12acc1ed644df1c3ee658068929c33863fccfdd2f8ab9236d58eb4496/cf2tf-0.8.0.tar.gz"
	cf2tf_cmd_tar := exec.Command("curl", "-L", cf2tf_tar_url, "-o", "package.tar.gz")
	err := cf2tf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cf2tf_zip_url := "https://files.pythonhosted.org/packages/52/00/94c12acc1ed644df1c3ee658068929c33863fccfdd2f8ab9236d58eb4496/cf2tf-0.8.0.zip"
	cf2tf_cmd_zip := exec.Command("curl", "-L", cf2tf_zip_url, "-o", "package.zip")
	err = cf2tf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cf2tf_bin_url := "https://files.pythonhosted.org/packages/52/00/94c12acc1ed644df1c3ee658068929c33863fccfdd2f8ab9236d58eb4496/cf2tf-0.8.0.bin"
	cf2tf_cmd_bin := exec.Command("curl", "-L", cf2tf_bin_url, "-o", "binary.bin")
	err = cf2tf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cf2tf_src_url := "https://files.pythonhosted.org/packages/52/00/94c12acc1ed644df1c3ee658068929c33863fccfdd2f8ab9236d58eb4496/cf2tf-0.8.0.src.tar.gz"
	cf2tf_cmd_src := exec.Command("curl", "-L", cf2tf_src_url, "-o", "source.tar.gz")
	err = cf2tf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cf2tf_cmd_direct := exec.Command("./binary")
	err = cf2tf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
