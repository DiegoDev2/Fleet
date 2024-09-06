package main

// KimApi - Knowledgebase of Interatomic Models (KIM) API
// Homepage: https://openkim.org

import (
	"fmt"
	
	"os/exec"
)

func installKimApi() {
	// Método 1: Descargar y extraer .tar.gz
	kimapi_tar_url := "https://s3.openkim.org/kim-api/kim-api-2.3.0.txz"
	kimapi_cmd_tar := exec.Command("curl", "-L", kimapi_tar_url, "-o", "package.tar.gz")
	err := kimapi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kimapi_zip_url := "https://s3.openkim.org/kim-api/kim-api-2.3.0.txz"
	kimapi_cmd_zip := exec.Command("curl", "-L", kimapi_zip_url, "-o", "package.zip")
	err = kimapi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kimapi_bin_url := "https://s3.openkim.org/kim-api/kim-api-2.3.0.txz"
	kimapi_cmd_bin := exec.Command("curl", "-L", kimapi_bin_url, "-o", "binary.bin")
	err = kimapi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kimapi_src_url := "https://s3.openkim.org/kim-api/kim-api-2.3.0.txz"
	kimapi_cmd_src := exec.Command("curl", "-L", kimapi_src_url, "-o", "source.tar.gz")
	err = kimapi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kimapi_cmd_direct := exec.Command("./binary")
	err = kimapi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
}
