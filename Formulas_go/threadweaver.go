package main

// Threadweaver - Helper for multithreaded programming
// Homepage: https://api.kde.org/frameworks/threadweaver/html/index.html

import (
	"fmt"
	
	"os/exec"
)

func installThreadweaver() {
	// Método 1: Descargar y extraer .tar.gz
	threadweaver_tar_url := "https://download.kde.org/stable/frameworks/6.5/threadweaver-6.5.0.tar.xz"
	threadweaver_cmd_tar := exec.Command("curl", "-L", threadweaver_tar_url, "-o", "package.tar.gz")
	err := threadweaver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	threadweaver_zip_url := "https://download.kde.org/stable/frameworks/6.5/threadweaver-6.5.0.tar.xz"
	threadweaver_cmd_zip := exec.Command("curl", "-L", threadweaver_zip_url, "-o", "package.zip")
	err = threadweaver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	threadweaver_bin_url := "https://download.kde.org/stable/frameworks/6.5/threadweaver-6.5.0.tar.xz"
	threadweaver_cmd_bin := exec.Command("curl", "-L", threadweaver_bin_url, "-o", "binary.bin")
	err = threadweaver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	threadweaver_src_url := "https://download.kde.org/stable/frameworks/6.5/threadweaver-6.5.0.tar.xz"
	threadweaver_cmd_src := exec.Command("curl", "-L", threadweaver_src_url, "-o", "source.tar.gz")
	err = threadweaver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	threadweaver_cmd_direct := exec.Command("./binary")
	err = threadweaver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: extra-cmake-modules")
exec.Command("latte", "install", "extra-cmake-modules")
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
}
