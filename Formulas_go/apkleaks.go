package main

// Apkleaks - Scanning APK file for URIs, endpoints & secrets
// Homepage: https://github.com/dwisiswant0/apkleaks

import (
	"fmt"
	
	"os/exec"
)

func installApkleaks() {
	// Método 1: Descargar y extraer .tar.gz
	apkleaks_tar_url := "https://files.pythonhosted.org/packages/40/88/8aa234dd5f7e632605dcce90d076982d4d1124d7278991ee54ec9e543cef/apkleaks-2.6.2.tar.gz"
	apkleaks_cmd_tar := exec.Command("curl", "-L", apkleaks_tar_url, "-o", "package.tar.gz")
	err := apkleaks_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apkleaks_zip_url := "https://files.pythonhosted.org/packages/40/88/8aa234dd5f7e632605dcce90d076982d4d1124d7278991ee54ec9e543cef/apkleaks-2.6.2.zip"
	apkleaks_cmd_zip := exec.Command("curl", "-L", apkleaks_zip_url, "-o", "package.zip")
	err = apkleaks_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apkleaks_bin_url := "https://files.pythonhosted.org/packages/40/88/8aa234dd5f7e632605dcce90d076982d4d1124d7278991ee54ec9e543cef/apkleaks-2.6.2.bin"
	apkleaks_cmd_bin := exec.Command("curl", "-L", apkleaks_bin_url, "-o", "binary.bin")
	err = apkleaks_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apkleaks_src_url := "https://files.pythonhosted.org/packages/40/88/8aa234dd5f7e632605dcce90d076982d4d1124d7278991ee54ec9e543cef/apkleaks-2.6.2.src.tar.gz"
	apkleaks_cmd_src := exec.Command("curl", "-L", apkleaks_src_url, "-o", "source.tar.gz")
	err = apkleaks_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apkleaks_cmd_direct := exec.Command("./binary")
	err = apkleaks_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jadx")
exec.Command("latte", "install", "jadx")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
