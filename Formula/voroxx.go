package main

// Voroxx - 3D Voronoi cell software library
// Homepage: https://math.lbl.gov/voro++

import (
	"fmt"
	
	"os/exec"
)

func installVoroxx() {
	// Método 1: Descargar y extraer .tar.gz
	voroxx_tar_url := "https://math.lbl.gov/voro++/download/dir/voro++-0.4.6.tar.gz"
	voroxx_cmd_tar := exec.Command("curl", "-L", voroxx_tar_url, "-o", "package.tar.gz")
	err := voroxx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	voroxx_zip_url := "https://math.lbl.gov/voro++/download/dir/voro++-0.4.6.zip"
	voroxx_cmd_zip := exec.Command("curl", "-L", voroxx_zip_url, "-o", "package.zip")
	err = voroxx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	voroxx_bin_url := "https://math.lbl.gov/voro++/download/dir/voro++-0.4.6.bin"
	voroxx_cmd_bin := exec.Command("curl", "-L", voroxx_bin_url, "-o", "binary.bin")
	err = voroxx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	voroxx_src_url := "https://math.lbl.gov/voro++/download/dir/voro++-0.4.6.src.tar.gz"
	voroxx_cmd_src := exec.Command("curl", "-L", voroxx_src_url, "-o", "source.tar.gz")
	err = voroxx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	voroxx_cmd_direct := exec.Command("./binary")
	err = voroxx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
