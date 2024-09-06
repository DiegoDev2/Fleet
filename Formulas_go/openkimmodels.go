package main

// OpenkimModels - All OpenKIM Models compatible with kim-api
// Homepage: https://openkim.org

import (
	"fmt"
	
	"os/exec"
)

func installOpenkimModels() {
	// Método 1: Descargar y extraer .tar.gz
	openkimmodels_tar_url := "https://s3.openkim.org/archives/collection/openkim-models-2021-08-11.txz"
	openkimmodels_cmd_tar := exec.Command("curl", "-L", openkimmodels_tar_url, "-o", "package.tar.gz")
	err := openkimmodels_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openkimmodels_zip_url := "https://s3.openkim.org/archives/collection/openkim-models-2021-08-11.txz"
	openkimmodels_cmd_zip := exec.Command("curl", "-L", openkimmodels_zip_url, "-o", "package.zip")
	err = openkimmodels_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openkimmodels_bin_url := "https://s3.openkim.org/archives/collection/openkim-models-2021-08-11.txz"
	openkimmodels_cmd_bin := exec.Command("curl", "-L", openkimmodels_bin_url, "-o", "binary.bin")
	err = openkimmodels_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openkimmodels_src_url := "https://s3.openkim.org/archives/collection/openkim-models-2021-08-11.txz"
	openkimmodels_cmd_src := exec.Command("curl", "-L", openkimmodels_src_url, "-o", "source.tar.gz")
	err = openkimmodels_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openkimmodels_cmd_direct := exec.Command("./binary")
	err = openkimmodels_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: kim-api")
exec.Command("latte", "install", "kim-api")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
}
