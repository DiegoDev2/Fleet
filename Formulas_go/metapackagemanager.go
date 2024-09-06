package main

// MetaPackageManager - Wrapper around all package managers with a unifying CLI
// Homepage: https://kdeldycke.github.io/meta-package-manager/

import (
	"fmt"
	
	"os/exec"
)

func installMetaPackageManager() {
	// Método 1: Descargar y extraer .tar.gz
	metapackagemanager_tar_url := "https://files.pythonhosted.org/packages/4f/7c/f31914ddedbd51616fb0e765bc1d25a8c4ec038a12c4f0b76206ce928443/meta_package_manager-5.18.0.tar.gz"
	metapackagemanager_cmd_tar := exec.Command("curl", "-L", metapackagemanager_tar_url, "-o", "package.tar.gz")
	err := metapackagemanager_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	metapackagemanager_zip_url := "https://files.pythonhosted.org/packages/4f/7c/f31914ddedbd51616fb0e765bc1d25a8c4ec038a12c4f0b76206ce928443/meta_package_manager-5.18.0.zip"
	metapackagemanager_cmd_zip := exec.Command("curl", "-L", metapackagemanager_zip_url, "-o", "package.zip")
	err = metapackagemanager_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	metapackagemanager_bin_url := "https://files.pythonhosted.org/packages/4f/7c/f31914ddedbd51616fb0e765bc1d25a8c4ec038a12c4f0b76206ce928443/meta_package_manager-5.18.0.bin"
	metapackagemanager_cmd_bin := exec.Command("curl", "-L", metapackagemanager_bin_url, "-o", "binary.bin")
	err = metapackagemanager_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	metapackagemanager_src_url := "https://files.pythonhosted.org/packages/4f/7c/f31914ddedbd51616fb0e765bc1d25a8c4ec038a12c4f0b76206ce928443/meta_package_manager-5.18.0.src.tar.gz"
	metapackagemanager_cmd_src := exec.Command("curl", "-L", metapackagemanager_src_url, "-o", "source.tar.gz")
	err = metapackagemanager_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	metapackagemanager_cmd_direct := exec.Command("./binary")
	err = metapackagemanager_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
