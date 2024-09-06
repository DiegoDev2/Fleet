package main

// EcsDeploy - CLI tool to simplify Amazon ECS deployments, rollbacks & scaling
// Homepage: https://github.com/fabfuel/ecs-deploy

import (
	"fmt"
	
	"os/exec"
)

func installEcsDeploy() {
	// Método 1: Descargar y extraer .tar.gz
	ecsdeploy_tar_url := "https://files.pythonhosted.org/packages/f1/3c/a2fc74f43992bda8df2e159351c254bacb5c157e766698b9aa537d459c7e/ecs-deploy-1.15.0.tar.gz"
	ecsdeploy_cmd_tar := exec.Command("curl", "-L", ecsdeploy_tar_url, "-o", "package.tar.gz")
	err := ecsdeploy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ecsdeploy_zip_url := "https://files.pythonhosted.org/packages/f1/3c/a2fc74f43992bda8df2e159351c254bacb5c157e766698b9aa537d459c7e/ecs-deploy-1.15.0.zip"
	ecsdeploy_cmd_zip := exec.Command("curl", "-L", ecsdeploy_zip_url, "-o", "package.zip")
	err = ecsdeploy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ecsdeploy_bin_url := "https://files.pythonhosted.org/packages/f1/3c/a2fc74f43992bda8df2e159351c254bacb5c157e766698b9aa537d459c7e/ecs-deploy-1.15.0.bin"
	ecsdeploy_cmd_bin := exec.Command("curl", "-L", ecsdeploy_bin_url, "-o", "binary.bin")
	err = ecsdeploy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ecsdeploy_src_url := "https://files.pythonhosted.org/packages/f1/3c/a2fc74f43992bda8df2e159351c254bacb5c157e766698b9aa537d459c7e/ecs-deploy-1.15.0.src.tar.gz"
	ecsdeploy_cmd_src := exec.Command("curl", "-L", ecsdeploy_src_url, "-o", "source.tar.gz")
	err = ecsdeploy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ecsdeploy_cmd_direct := exec.Command("./binary")
	err = ecsdeploy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
