package main

// Tika - Content analysis toolkit
// Homepage: https://tika.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installTika() {
	// Método 1: Descargar y extraer .tar.gz
	tika_tar_url := "https://www.apache.org/dyn/closer.lua?path=tika/2.9.2/tika-app-2.9.2.jar"
	tika_cmd_tar := exec.Command("curl", "-L", tika_tar_url, "-o", "package.tar.gz")
	err := tika_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tika_zip_url := "https://www.apache.org/dyn/closer.lua?path=tika/2.9.2/tika-app-2.9.2.jar"
	tika_cmd_zip := exec.Command("curl", "-L", tika_zip_url, "-o", "package.zip")
	err = tika_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tika_bin_url := "https://www.apache.org/dyn/closer.lua?path=tika/2.9.2/tika-app-2.9.2.jar"
	tika_cmd_bin := exec.Command("curl", "-L", tika_bin_url, "-o", "binary.bin")
	err = tika_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tika_src_url := "https://www.apache.org/dyn/closer.lua?path=tika/2.9.2/tika-app-2.9.2.jar"
	tika_cmd_src := exec.Command("curl", "-L", tika_src_url, "-o", "source.tar.gz")
	err = tika_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tika_cmd_direct := exec.Command("./binary")
	err = tika_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
