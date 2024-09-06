package main

// Librttopo - RT Topology Library
// Homepage: https://git.osgeo.org/gitea/rttopo/librttopo

import (
	"fmt"
	
	"os/exec"
)

func installLibrttopo() {
	// Método 1: Descargar y extraer .tar.gz
	librttopo_tar_url := "https://git.osgeo.org/gitea/rttopo/librttopo/archive/librttopo-1.1.0.tar.gz"
	librttopo_cmd_tar := exec.Command("curl", "-L", librttopo_tar_url, "-o", "package.tar.gz")
	err := librttopo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	librttopo_zip_url := "https://git.osgeo.org/gitea/rttopo/librttopo/archive/librttopo-1.1.0.zip"
	librttopo_cmd_zip := exec.Command("curl", "-L", librttopo_zip_url, "-o", "package.zip")
	err = librttopo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	librttopo_bin_url := "https://git.osgeo.org/gitea/rttopo/librttopo/archive/librttopo-1.1.0.bin"
	librttopo_cmd_bin := exec.Command("curl", "-L", librttopo_bin_url, "-o", "binary.bin")
	err = librttopo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	librttopo_src_url := "https://git.osgeo.org/gitea/rttopo/librttopo/archive/librttopo-1.1.0.src.tar.gz"
	librttopo_cmd_src := exec.Command("curl", "-L", librttopo_src_url, "-o", "source.tar.gz")
	err = librttopo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	librttopo_cmd_direct := exec.Command("./binary")
	err = librttopo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: geos")
exec.Command("latte", "install", "geos")
}
