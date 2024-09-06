package main

// Dvc - Git for data science projects
// Homepage: https://dvc.org

import (
	"fmt"
	
	"os/exec"
)

func installDvc() {
	// Método 1: Descargar y extraer .tar.gz
	dvc_tar_url := "https://files.pythonhosted.org/packages/ce/fe/30d8a9733f1edddd356766f0ec167e950e26a4717ec6f0c397da25f9510f/dvc-3.55.2.tar.gz"
	dvc_cmd_tar := exec.Command("curl", "-L", dvc_tar_url, "-o", "package.tar.gz")
	err := dvc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dvc_zip_url := "https://files.pythonhosted.org/packages/ce/fe/30d8a9733f1edddd356766f0ec167e950e26a4717ec6f0c397da25f9510f/dvc-3.55.2.zip"
	dvc_cmd_zip := exec.Command("curl", "-L", dvc_zip_url, "-o", "package.zip")
	err = dvc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dvc_bin_url := "https://files.pythonhosted.org/packages/ce/fe/30d8a9733f1edddd356766f0ec167e950e26a4717ec6f0c397da25f9510f/dvc-3.55.2.bin"
	dvc_cmd_bin := exec.Command("curl", "-L", dvc_bin_url, "-o", "binary.bin")
	err = dvc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dvc_src_url := "https://files.pythonhosted.org/packages/ce/fe/30d8a9733f1edddd356766f0ec167e950e26a4717ec6f0c397da25f9510f/dvc-3.55.2.src.tar.gz"
	dvc_cmd_src := exec.Command("curl", "-L", dvc_src_url, "-o", "source.tar.gz")
	err = dvc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dvc_cmd_direct := exec.Command("./binary")
	err = dvc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: apache-arrow")
	exec.Command("latte", "install", "apache-arrow").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: libgit2")
	exec.Command("latte", "install", "libgit2").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: numpy")
	exec.Command("latte", "install", "numpy").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: patchelf")
	exec.Command("latte", "install", "patchelf").Run()
}
