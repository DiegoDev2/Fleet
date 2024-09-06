package main

// Ki18n - KDE Gettext-based UI text internationalization
// Homepage: https://api.kde.org/frameworks/ki18n/html/index.html

import (
	"fmt"
	
	"os/exec"
)

func installKi18n() {
	// Método 1: Descargar y extraer .tar.gz
	ki18n_tar_url := "https://download.kde.org/stable/frameworks/6.5/ki18n-6.5.0.tar.xz"
	ki18n_cmd_tar := exec.Command("curl", "-L", ki18n_tar_url, "-o", "package.tar.gz")
	err := ki18n_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ki18n_zip_url := "https://download.kde.org/stable/frameworks/6.5/ki18n-6.5.0.tar.xz"
	ki18n_cmd_zip := exec.Command("curl", "-L", ki18n_zip_url, "-o", "package.zip")
	err = ki18n_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ki18n_bin_url := "https://download.kde.org/stable/frameworks/6.5/ki18n-6.5.0.tar.xz"
	ki18n_cmd_bin := exec.Command("curl", "-L", ki18n_bin_url, "-o", "binary.bin")
	err = ki18n_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ki18n_src_url := "https://download.kde.org/stable/frameworks/6.5/ki18n-6.5.0.tar.xz"
	ki18n_cmd_src := exec.Command("curl", "-L", ki18n_src_url, "-o", "source.tar.gz")
	err = ki18n_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ki18n_cmd_direct := exec.Command("./binary")
	err = ki18n_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
	fmt.Println("Instalando dependencia: extra-cmake-modules")
	exec.Command("latte", "install", "extra-cmake-modules").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: iso-codes")
	exec.Command("latte", "install", "iso-codes").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
}
