package main

// Openfa - Set of algorithms that implement standard models used in fundamental astronomy
// Homepage: https://gitlab.obspm.fr/imcce_openfa/openfa

import (
	"fmt"
	
	"os/exec"
)

func installOpenfa() {
	// Método 1: Descargar y extraer .tar.gz
	openfa_tar_url := "https://gitlab.obspm.fr/imcce_openfa/openfa/-/archive/20231011.0.3/openfa-20231011.0.3.tar.gz"
	openfa_cmd_tar := exec.Command("curl", "-L", openfa_tar_url, "-o", "package.tar.gz")
	err := openfa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openfa_zip_url := "https://gitlab.obspm.fr/imcce_openfa/openfa/-/archive/20231011.0.3/openfa-20231011.0.3.zip"
	openfa_cmd_zip := exec.Command("curl", "-L", openfa_zip_url, "-o", "package.zip")
	err = openfa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openfa_bin_url := "https://gitlab.obspm.fr/imcce_openfa/openfa/-/archive/20231011.0.3/openfa-20231011.0.3.bin"
	openfa_cmd_bin := exec.Command("curl", "-L", openfa_bin_url, "-o", "binary.bin")
	err = openfa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openfa_src_url := "https://gitlab.obspm.fr/imcce_openfa/openfa/-/archive/20231011.0.3/openfa-20231011.0.3.src.tar.gz"
	openfa_cmd_src := exec.Command("curl", "-L", openfa_src_url, "-o", "source.tar.gz")
	err = openfa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openfa_cmd_direct := exec.Command("./binary")
	err = openfa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
}
