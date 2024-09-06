package main

// HopscotchMap - C++ implementation of a fast hash map and hash set using hopscotch hashing
// Homepage: https://github.com/Tessil/hopscotch-map

import (
	"fmt"
	
	"os/exec"
)

func installHopscotchMap() {
	// Método 1: Descargar y extraer .tar.gz
	hopscotchmap_tar_url := "https://github.com/Tessil/hopscotch-map/archive/refs/tags/v2.3.1.tar.gz"
	hopscotchmap_cmd_tar := exec.Command("curl", "-L", hopscotchmap_tar_url, "-o", "package.tar.gz")
	err := hopscotchmap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hopscotchmap_zip_url := "https://github.com/Tessil/hopscotch-map/archive/refs/tags/v2.3.1.zip"
	hopscotchmap_cmd_zip := exec.Command("curl", "-L", hopscotchmap_zip_url, "-o", "package.zip")
	err = hopscotchmap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hopscotchmap_bin_url := "https://github.com/Tessil/hopscotch-map/archive/refs/tags/v2.3.1.bin"
	hopscotchmap_cmd_bin := exec.Command("curl", "-L", hopscotchmap_bin_url, "-o", "binary.bin")
	err = hopscotchmap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hopscotchmap_src_url := "https://github.com/Tessil/hopscotch-map/archive/refs/tags/v2.3.1.src.tar.gz"
	hopscotchmap_cmd_src := exec.Command("curl", "-L", hopscotchmap_src_url, "-o", "source.tar.gz")
	err = hopscotchmap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hopscotchmap_cmd_direct := exec.Command("./binary")
	err = hopscotchmap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
