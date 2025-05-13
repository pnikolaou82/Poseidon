
# 🌊 Poseidon

**Poseidon** is an internal Taikun platform that combines **Kubernetes Operators** (built with Operator SDK) and a **Flask-based dashboard** to bring SRE excellence to modern cloud-native environments. Named after the Greek god of the sea, Poseidon offers control over the chaos of infrastructure just as the deity ruled the oceans.

---

## ⚙️ Overview

Poseidon provides:

- 📦 Custom Kubernetes **Operators** to automate SRE tasks like log cleanup, alert tuning, service healing, and configuration reconciliation.
- 📊 A lightweight **Flask Dashboard** to monitor metrics, visualize operator activity, and inspect anomalies.
- 🌐 Seamless integration with observability stacks like **Prometheus**, **Alertmanager**, **Grafana**, and **Loki**.
- 💡 GitOps compatibility for declaring SRE policies and configurations.

---

## 🏗️ Architecture

```
 Kubernetes Cluster
 ┌───────────────────────────┐
 │        Poseidon           │
 │ ┌─────────────┐           │
 │ │ Operators   │           │
 │ │ (Go/Ansible)│ <-- CRDs  │
 │ └─────────────┘           │
 │       ⬇                   │
 │ ┌──────────────┐          │
 │ │ Flask SRE UI │ <--> Prometheus / Grafana / Loki
 │ └──────────────┘          │
 └───────────────────────────┘
```

---

## 📁 Repository Structure

```
poseidon/
├── operators/                # Operator SDK projects (Go or Ansible)
│   └── prometheus-tuner/    # Example operator for tuning alert thresholds
│
├── dashboard/               # Flask-based SRE dashboard
│   ├── app/
│   ├── templates/
│   ├── static/
│   └── ...
│
├── helm-charts/             # Helm chart for deploying Poseidon
│
└── README.md
```

---

## ✨ Key Features

| Feature                        | Description |
|-------------------------------|-------------|
| **Operator Automation**       | Custom logic for monitoring and infrastructure health |
| **Dashboard UI**              | View live metrics, operator activity, and health checks |
| **Anomaly Detection Hooks**   | Optional hooks for anomaly alerting and triage |
| **RBAC & Namespacing**        | Multi-tenant control over operator scope |
| **GitOps Compatibility**      | Watch Git for CR definitions |
| **Historical Logs & Charts**  | Time-series views of previous remediations |

---

## 🚀 Getting Started

### Prerequisites

- Kubernetes Cluster (e.g. minikube, GKE, EKS)
- Python 3.9+
- [Operator SDK](https://sdk.operatorframework.io/)
- Prometheus & Grafana (optional)

### Clone & Run Dashboard

```bash
git clone git@github.com:taikun-cloud/poseidon.git
cd poseidon/dashboard
pip install -r requirements.txt
export FLASK_APP=app
flask run
```

### Deploy an Operator

```bash
cd poseidon/operators/prometheus-tuner
make install
make run
```

---

## 🧭 Use Cases

- Automatically reconcile ResourceRule CRs
- Auto-tune alert thresholds based on usage patterns
- Visualize noisy alerts and their trends
- Schedule and audit cleanup of expired Kubernetes resources
- Run hooks for restart loops or pod crash detection

---

## 📈 Roadmap

- [ ] Add AI-assisted root cause suggestions to dashboard
- [ ] Integrate Loki-based log summaries into Flask UI
- [ ] Extend operators for config drift detection
- [ ] Add Slack / Webhook notifications

---

## 🧙 Mythological Context

The name **Poseidon** reflects Kubernetes’ Greek etymology ("κυβερνήτης" = *helmsman*) and the project’s goal: to navigate, steer, and command the turbulent seas of cloud infrastructure.

---

## 🏛️ License

This project is proprietary and maintained internally by the **Taikun SRE team**.

---

## 🙏 Acknowledgments

- [Operator SDK](https://sdk.operatorframework.io/)
- [Flask](https://flask.palletsprojects.com/)
- [Prometheus](https://prometheus.io/)
- [Grafana](https://grafana.com/)
