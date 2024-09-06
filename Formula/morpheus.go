package main

// Morpheus - Modeling environment for multi-cellular systems biology
// Homepage: https://morpheus.gitlab.io/

import (
	"fmt"
	
	"os/exec"
)

func installMorpheus() {
	// Método 1: Descargar y extraer .tar.gz
	morpheus_tar_url := "https://gitlab.com/morpheus.lab/morpheus/-/archive/v2.3.8/morpheus-v2.3.8.tar.gz"
	morpheus_cmd_tar := exec.Command("curl", "-L", morpheus_tar_url, "-o", "package.tar.gz")
	err := morpheus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	morpheus_zip_url := "https://gitlab.com/morpheus.lab/morpheus/-/archive/v2.3.8/morpheus-v2.3.8.zip"
	morpheus_cmd_zip := exec.Command("curl", "-L", morpheus_zip_url, "-o", "package.zip")
	err = morpheus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	morpheus_bin_url := "https://gitlab.com/morpheus.lab/morpheus/-/archive/v2.3.8/morpheus-v2.3.8.bin"
	morpheus_cmd_bin := exec.Command("curl", "-L", morpheus_bin_url, "-o", "binary.bin")
	err = morpheus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	morpheus_src_url := "https://gitlab.com/morpheus.lab/morpheus/-/archive/v2.3.8/morpheus-v2.3.8.src.tar.gz"
	morpheus_cmd_src := exec.Command("curl", "-L", morpheus_src_url, "-o", "source.tar.gz")
	err = morpheus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	morpheus_cmd_direct := exec.Command("./binary")
	err = morpheus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: graphviz")
	exec.Command("latte", "install", "graphviz").Run()
	fmt.Println("Instalando dependencia: libomp")
	exec.Command("latte", "install", "libomp").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: qt@5")
	exec.Command("latte", "install", "qt@5").Run()
}
