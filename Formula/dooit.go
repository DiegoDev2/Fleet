package main

// Dooit - TUI todo manager
// Homepage: https://github.com/kraanzu/dooit

import (
	"fmt"
	
	"os/exec"
)

func installDooit() {
	// Método 1: Descargar y extraer .tar.gz
	dooit_tar_url := "https://files.pythonhosted.org/packages/eb/f3/3d12b29b9ec6302197b215890fe9b0263f3f78f148270513ae4eb7b89f77/dooit-2.2.0.tar.gz"
	dooit_cmd_tar := exec.Command("curl", "-L", dooit_tar_url, "-o", "package.tar.gz")
	err := dooit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dooit_zip_url := "https://files.pythonhosted.org/packages/eb/f3/3d12b29b9ec6302197b215890fe9b0263f3f78f148270513ae4eb7b89f77/dooit-2.2.0.zip"
	dooit_cmd_zip := exec.Command("curl", "-L", dooit_zip_url, "-o", "package.zip")
	err = dooit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dooit_bin_url := "https://files.pythonhosted.org/packages/eb/f3/3d12b29b9ec6302197b215890fe9b0263f3f78f148270513ae4eb7b89f77/dooit-2.2.0.bin"
	dooit_cmd_bin := exec.Command("curl", "-L", dooit_bin_url, "-o", "binary.bin")
	err = dooit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dooit_src_url := "https://files.pythonhosted.org/packages/eb/f3/3d12b29b9ec6302197b215890fe9b0263f3f78f148270513ae4eb7b89f77/dooit-2.2.0.src.tar.gz"
	dooit_cmd_src := exec.Command("curl", "-L", dooit_src_url, "-o", "source.tar.gz")
	err = dooit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dooit_cmd_direct := exec.Command("./binary")
	err = dooit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
